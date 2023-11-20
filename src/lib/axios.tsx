import axios from "axios";

const sucessfullStatusCode = [200, 400];

const client = axios.create({
	baseURL: "http://localhost:8000",
	validateStatus: (status) => sucessfullStatusCode.includes(status),
	withCredentials: true,
});

client.defaults.headers.common["Content-Type"] = "application/json";

export default client;
