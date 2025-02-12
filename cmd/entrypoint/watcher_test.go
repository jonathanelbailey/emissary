package entrypoint_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datawire/dlib/dlog"
	"github.com/emissary-ingress/emissary/v3/cmd/entrypoint"
	"github.com/emissary-ingress/emissary/v3/pkg/kates"
	"github.com/emissary-ingress/emissary/v3/pkg/snapshot/v1"
)

func TestGetDeltaType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		op       string
		expected kates.DeltaType
	}{
		{
			name:     "Add operation",
			op:       "add",
			expected: kates.ObjectAdd,
		},
		{
			name:     "Update operation",
			op:       "update",
			expected: kates.ObjectUpdate,
		},
		{
			name:     "Delete operation",
			op:       "delete",
			expected: kates.ObjectDelete,
		},
		{
			name:     "Unknown operation",
			op:       "unknown",
			expected: kates.ObjectAdd,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := entrypoint.GetDeltaType(tt.op)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestIstioUpdate(t *testing.T) {
	t.Parallel()

	// given
	istioCertName := "istio-certs"
	istioCertNamespace := "ambassador"

	// when
	tests := []struct {
		name          string
		icertUpdate   entrypoint.IstioCertUpdate
		expectedDelta *kates.Delta
	}{
		{
			name: "Add new certificate",
			icertUpdate: entrypoint.IstioCertUpdate{
				Op:        "add",
				Name:      istioCertName,
				Namespace: istioCertNamespace,
				Secret:    nil,
			},
			expectedDelta: &kates.Delta{
				TypeMeta: kates.TypeMeta{
					APIVersion: "v1",
					Kind:       "Secret",
				},
				ObjectMeta: kates.ObjectMeta{
					Name:      istioCertName,
					Namespace: istioCertNamespace,
				},
				DeltaType: kates.ObjectAdd,
			},
		},
		{
			name: "Update existing certificate",
			icertUpdate: entrypoint.IstioCertUpdate{
				Op:        "update",
				Name:      istioCertName,
				Namespace: istioCertNamespace,
				Secret:    nil,
			},
			expectedDelta: &kates.Delta{
				TypeMeta: kates.TypeMeta{
					APIVersion: "v1",
					Kind:       "Secret",
				},
				ObjectMeta: kates.ObjectMeta{
					Name:      istioCertName,
					Namespace: istioCertNamespace,
				},
				DeltaType: kates.ObjectUpdate,
			},
		},
		{
			name: "Delete certificate",
			icertUpdate: entrypoint.IstioCertUpdate{
				Op:        "delete",
				Name:      istioCertName,
				Namespace: istioCertNamespace,
				Secret:    nil,
			},
			expectedDelta: &kates.Delta{
				TypeMeta: kates.TypeMeta{
					APIVersion: "v1",
					Kind:       "Secret",
				},
				ObjectMeta: kates.ObjectMeta{
					Name:      istioCertName,
					Namespace: istioCertNamespace,
				},
				DeltaType: kates.ObjectDelete,
			},
		},
	}

	// then
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := dlog.NewTestContext(t, false)

			// Create snapshot holder
			sh, err := entrypoint.NewSnapshotHolder(&snapshot.AmbassadorMetaInfo{})
			require.NoError(t, err)
			istioCertSrc := entrypoint.NewIstioCertSource()
			istioCertWatcher, err := istioCertSrc.Watch(ctx)
			require.NoError(t, err)

			istio := entrypoint.NewIstioCertWatchManager(ctx, istioCertWatcher)

			// Call IstioUpdate
			changed, err := sh.IstioUpdate(ctx, istio, tt.icertUpdate)
			require.NoError(t, err)
			require.True(t, changed)

			// Verify delta was created and added to unsentDeltas
			require.Len(t, sh.GetUnsentDeltas(), 1)
			delta := sh.GetUnsentDeltas()[0]

			require.Equal(t, tt.expectedDelta.APIVersion, delta.APIVersion)
			require.Equal(t, tt.expectedDelta.Kind, delta.Kind)
			require.Equal(t, tt.expectedDelta.Name, delta.Name)
			require.Equal(t, tt.expectedDelta.Namespace, delta.Namespace)
			require.Equal(t, tt.expectedDelta.DeltaType, delta.DeltaType)
		})
	}
}
