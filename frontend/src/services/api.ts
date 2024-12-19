import axios from "axios";
import { getDefaultStore } from "jotai/vanilla";
import { authTokenAtom } from "@/store/token";

const apiClient = axios.create({
  baseURL: "http://localhost:8000",
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

apiClient.interceptors.request.use((config) => {
  const token = getDefaultStore().get(authTokenAtom); // Access token from Jotai
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default apiClient;
