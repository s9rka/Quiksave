import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import "./App.css";
import LoginForm from "./components/auth/Login";
import { Routes, Route, useLocation, Navigate } from "react-router-dom";
import Register from "./components/auth/Register";
import Dashboard from "./components/Dashboard";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Routes>
        <Route path="/" element={<LoginForm />} />
        <Route path="/register" element={<Register/>} />

        <Route path="/:username" element={<Dashboard/>} />
        
        
      </Routes>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;
