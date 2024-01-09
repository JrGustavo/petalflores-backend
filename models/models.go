package models

type SecretRDSJaon struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Egine               string `json:"engine"`
	Host                string `json:"host"`
	Port                string `json:"port"`
	DbClusterIdentifier string `json:"host"`
}

type Signup struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
