package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {

	//如果取不到当前用户的家目录，就没办法设置kubeconfig的默认目录了，只能从入参中取
	kubeconfig := "xxx-config"

	//从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("load kubeconfig failed!err：", err)
		panic(err.Error())
	}

	//实例化一个clientset对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("init clientset failed ! err: ", err)
		panic(err.Error())
	}

	//      if obj.Spec.Paused {
	//			return nil, errors.New("can't restart paused deployment (run rollout resume first)")
	//		}
	//		if obj.Spec.Template.ObjectMeta.Annotations == nil {
	//			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	//		}
	//		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	klog.Info(formatTimeStr)
	content := fmt.Sprintf("%s%s%s", "{\"spec\": {\"template\": {\"metadata\":{\"annotations\":{\"kubectl.kubernetes.io/restartedAt\": \"", formatTimeStr, "\"}}}}}")
	_, err = clientset.AppsV1().Deployments("infra").Patch(context.TODO(), "xxx-v1", types.StrategicMergePatchType, []byte(content), metav1.PatchOptions{})
	if err != nil {
		klog.Error(err)
	}

}
