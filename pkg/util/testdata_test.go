/*
Copyright 2020-2021 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"time"

	cstortypes "github.com/openebs/api/v2/pkg/apis/types"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	fourGigiByte    = resource.MustParse("4Gi")
	cstorScName     = "cstor-sc"
	cstorVolumeMode = corev1.PersistentVolumeFilesystem

	cstorPV1 = corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "pvc-1",
			CreationTimestamp: metav1.Time{Time: time.Now()},
			Labels:            map[string]string{cstortypes.PersistentVolumeLabelKey: "pvc-1"},
			Finalizers:        []string{},
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity:    corev1.ResourceList{corev1.ResourceStorage: fourGigiByte},
			AccessModes: []corev1.PersistentVolumeAccessMode{"ReadWriteOnce"},
			ClaimRef: &corev1.ObjectReference{
				Namespace: "default",
				Name:      "cstor-pvc-1",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			StorageClassName:              cstorScName,
			VolumeMode:                    &cstorVolumeMode,
			PersistentVolumeSource: corev1.PersistentVolumeSource{CSI: &corev1.CSIPersistentVolumeSource{
				Driver: "cstor.csi.openebs.io", VolumeAttributes: map[string]string{"openebs.io/cas-type": "cstor"},
			}},
		},
		Status: corev1.PersistentVolumeStatus{Phase: corev1.VolumeBound},
	}
	zfspv = corev1.PersistentVolume{
		Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeSource: corev1.PersistentVolumeSource{CSI: &corev1.CSIPersistentVolumeSource{Driver: ZFSCSIDriver}}}}

	cstorSC = v1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "cstor-sc",
			CreationTimestamp: metav1.Time{Time: time.Now()},
		},
		Provisioner: "cstor.csi.openebs.io",
	}

	jivaSC = v1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "jiva-sc",
			CreationTimestamp: metav1.Time{Time: time.Now()},
		},
		Provisioner: "jiva.csi.openebs.io",
		Parameters:  map[string]string{"cas-type": "jiva"},
	}
)
