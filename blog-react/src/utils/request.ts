import axios, {
  type AxiosRequestConfig,
  type AxiosResponse,
  type InternalAxiosRequestConfig,
} from "axios";

const service = axios.create({
  baseURL: "/api",
  timeout: 10000,
});

export interface ApiResponse<T> {
  code: number;
  msg: string;
  data: T;
}

service.interceptors.request.use((config: AxiosRequestConfig) => {
  const token = localStorage.getItem("access_token");
  config.headers = {
    "Content-Type": "application/json",
    "x-access-token": token,
    ...config.headers,
  };
  return config as InternalAxiosRequestConfig;
});

service.interceptors.response.use((response: AxiosResponse) => {
  return response.data;
});

export default service;
