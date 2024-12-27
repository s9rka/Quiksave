import axios from "axios";

export const privateClient = axios.create({
  baseURL: "https://localhost:8443",
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

export const publicClient = axios.create({
  baseURL: "https://localhost:8443",
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});
