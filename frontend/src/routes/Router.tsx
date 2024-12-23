import { Routes, Route } from "react-router-dom";
import LoginForm from "../components/auth/Login";
import Register from "../components/auth/Register";
import Dashboard from "../components/storage/Storage";
import { ProtectedRoute } from "./ProtectedRoute";
import CreateForm from "@/components/createNote/CreateForm";
import {PublicRoute} from "./PublicRoute";

const AppRouter = () => {
  return (
    <Routes>
      <Route element={<PublicRoute />}>
        <Route path="/" element={<LoginForm />} />
        <Route path="/register" element={<Register />} />
      </Route>

      <Route element={<ProtectedRoute />}>
        <Route path="/:username" element={<Dashboard />} />
        <Route path="/new" element={<CreateForm />} />
      </Route>
    </Routes>
  );
};

export default AppRouter;
