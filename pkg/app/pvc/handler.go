package pvc

import (
	"client-go-demo/pkg/middleware"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"

	. "client-go-demo/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 获取pvc列表
func listPersistentVolumeClaim(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Param("namespace")
	labelSelector := c.Query("labelSelector")

	persistentVolumeClaimList, err := client.CoreV1().PersistentVolumeClaims(namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pvc list failed, error:%v", err.Error())
		c.JSON(400, "get pvc list failed")
		return
	}
	persistentVolumeClaimList.Kind = "List"
	persistentVolumeClaimList.APIVersion = "v1"

	for index := range persistentVolumeClaimList.Items {
		persistentVolumeClaimList.Items[index].APIVersion = "v1"
		persistentVolumeClaimList.Items[index].Kind = "PersistentVolumeClaim"
	}
	c.JSON(200, persistentVolumeClaimList)
	return
}

// 获取pvc详情
func getPersistentVolumeClaim(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Param("namespace")
	name := c.Param("name")

	persistentVolumeClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pvc failed, error:%v", err.Error())
		c.JSON(500, "get pvc failed")
		return
	}

	persistentVolumeClaim.APIVersion = "v1"
	persistentVolumeClaim.Kind = "PersistentVolumeClaim"

	c.JSON(200, persistentVolumeClaim)

}

// 创建pvc
func createPersistentVolumeClaim(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Param("namespace")

	pvcRequest := &corev1.PersistentVolumeClaim{}
	err := c.BindJSON(pvcRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "gett create pvc body failed, error:%v", err.Error())
		c.JSON(400, err)
		return
	}

	persistentVolumeClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Create(pvcRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "create pvc failed, error:%v", err.Error())
		c.JSON(400, err)
		return
	}

	persistentVolumeClaim.APIVersion = "v1"
	persistentVolumeClaim.Kind = "PersistentVolumeClaim"

	c.JSON(200, persistentVolumeClaim)

}

// 更新pvc
func updatePersistentVolumeClaim(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Param("namespace")
	name := c.Param("name")

	// 序列化请求体信息
	pvcRequest := &corev1.PersistentVolumeClaim{}
	err := c.BindJSON(pvcRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get update pvc body failed, error:%v", err.Error())
		c.JSON(400, err)
		return
	}

	// 根据名称获取pvc,写入ResourceVersion
	persistentVolumeClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "update pvc failed, error:%v", err.Error())
		c.JSON(400, err)
		return
	}
	pvcRequest.ResourceVersion = persistentVolumeClaim.ResourceVersion

	// 执行更新
	updatePersistentVolumeClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Update(pvcRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "update pvc failed, error:%v", err.Error())
		c.JSON(400, err)
		return
	}
	c.JSON(200, updatePersistentVolumeClaim)
}

// 删除PVC
func deletePersistentVolumeClaim(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Param("namespace")
	name := c.Param("name")

	err := client.CoreV1().PersistentVolumeClaims(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "delete pvc failed, error:%v", err.Error())
		c.JSON(400, err)
		return
	}

	c.JSON(200, struct{}{})

}
