import { Routes, Route } from "react-router-dom";
import LoginForm from "../components/auth/Login";
import Register from "../components/auth/Register";
import Dashboard from "../components/storage/Storage";
import {PrivateRoutes} from "./PrivateRoute";
import {PublicRoutes} from "./PublicRoute";

const AppRouter = () => {
  return (
    <Routes>
      <Route element={<PublicRoutes />}>
        <Route path="/" element={<LoginForm />} />
        <Route path="/register" element={<Register />} />
      </Route>

      <Route element={<PrivateRoutes />}>
        <Route path="/:username" element={<Dashboard />} />
      </Route>
    </Routes>
  );
}

export default AppRouter;
