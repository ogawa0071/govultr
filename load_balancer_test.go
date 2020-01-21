package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLoadBalancerHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `[{"SUBID":1317575,"date_created":"2020-01-07 17:24:23","location":"New Jersey","label":"test","status":"active"}]`
		fmt.Fprintf(writer, response)
	})

	list, err := client.LoadBalancer.List(ctx)

	if err != nil {
		t.Errorf("LoadBalancer.List returned %+v, ", err)
	}

	expected := []LoadBalancers{
		{
			ID:          1317575,
			DateCreated: "2020-01-07 17:24:23",
			Location:    "New Jersey",
			Label:       "test",
			Status:      "active",
			RegionID:    0,
			IPV6:        "",
			IPV4:        "",
		},
	}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("LoadBalancer.List returned %+v, expected %+v", list, expected)
	}
}

func TestLoadBalancerHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.Delete(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.Delete returned %+v, ", err)
	}
}

func TestLoadBalancerHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.SetLabel(ctx, 12345, "label")

	if err != nil {
		t.Errorf("LoadBalancer.SetLabel returned %+v, ", err)
	}
}

func TestLoadBalancerHandler_AttachedInstances(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/instance_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"instance_list": ["1234", "2341"]}`
		fmt.Fprintf(writer, response)
	})

	instanceList, err := client.LoadBalancer.AttachedInstances(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.AttachedInstances returned %+v, ", err)
	}

	expected := &InstanceList{InstanceList: []string{"1234", "2341"}}

	if !reflect.DeepEqual(instanceList, expected) {
		t.Errorf("LoadBalancer.AttachedInstances returned %+v, expected %+v", instanceList, expected)
	}
}

func TestLoadBalancerHandler_AttachInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/instance_attach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.AttachInstance(ctx, 12345, 45678)

	if err != nil {
		t.Errorf("LoadBalancer.AttachInstance returned %+v, ", err)
	}
}

func TestLoadBalancerHandler_DetachInstance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/instance_detach", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.LoadBalancer.DetachInstance(ctx, 12345, 45678)

	if err != nil {
		t.Errorf("LoadBalancer.DetachInstance returned %+v, ", err)
	}
}

func TestLoadBalancerHandler_GetHealthCheck(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/health_check_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "protocol": "http","port": 81,"path": "/test","check_interval": 10,"response_timeout": 45,"unhealthy_threshold": 1,"healthy_threshold": 2}`
		fmt.Fprintf(writer, response)
	})

	health, err := client.LoadBalancer.GetHealthCheck(ctx, 12345)

	if err != nil {
		t.Errorf("LoadBalancer.GetHealthCheck returned %+v, ", err)
	}

	expected := &HealthCheck{
		Protocol:           "http",
		Port:               81,
		Path:               "/test",
		CheckInterval:      10,
		ResponseTimeout:    45,
		UnhealthyThreshold: 1,
		HealthyThreshold:   2,
	}

	if !reflect.DeepEqual(health, expected) {
		t.Errorf("LoadBalancer.GetHealthCheck returned %+v, expected %+v", health, expected)
	}
}

func TestLoadBalancerHandler_SetHealthCheck(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancer/health_check_update", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	health := &HealthCheck{
		Protocol:           "HTTPS",
		Port:               8080,
		Path:               "/health",
		CheckInterval:      4,
		ResponseTimeout:    5,
		UnhealthyThreshold: 3,
		HealthyThreshold:   4,
	}
	err := client.LoadBalancer.SetHealthCheck(ctx, 12345, health)

	if err != nil {
		t.Errorf("LoadBalancer.SetHealthCheck returned %+v, ", err)
	}
}
