import axios from "axios";
import { getToken } from "@/services/auth";

export const privateClient = axios.create({
  baseURL: "http://localhost:8000",
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

privateClient.interceptors.request.use((config) => {
  const token = getToken();
  console.log("Auth Token:", token);

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const publicClient = axios.create({
  baseURL: "http://localhost:8000",
  headers: {
    "Content-Type": "application/json",
  },
});
