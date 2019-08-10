package sim

import "time"




var (
	sim_timestamp int64
	sim_time time.Time
	sim_set = false

	sim_running = false

)

func Start(){
	sim_running = true
}

func Stop() {
	sim_running = false
}

func IsRunning() bool {
	return sim_running
}


func SetSimulationTime(t int64){
	sim_timestamp = t
	sim_time = time.Unix(t,0)
	sim_set = true
}

func DisableSimulationTime(){
	sim_set = false
}


func Unix() int64{
	if sim_set{
		return sim_timestamp
	}
	return time.Now().Unix()
}

func UTC() time.Time{
	if sim_set{
		return sim_time
	}
	return time.Now().UTC()
}