import { Routes, Route } from "react-router-dom";
import LoginForm from "../components/auth/Login";
import Register from "../components/auth/Register";
import Storage from "../components/storage/Storage";
import CreateNote from "@/components/createNote/CreateNote";
import { PublicRoute } from "./PublicRoute";
import ProtectedRoute from "./ProtectedRoute";
import AppLayout from "@/components/layout/AppLayout";
import { EditNotePage } from "@/components/createNote/EditNote";
import Logout from "@/components/auth/Logout";
import CreateVault from "@/components/CreateVault";
import VaultList from "@/components/VaultList";

const AppRouter = () => {
  return (
    <Routes>
      <Route element={<AppLayout />}>
        {/* Public Routes */}
        <Route element={<PublicRoute />}>
          <Route path="/" element={<LoginForm />} />
          <Route path="/register" element={<Register />} />
        </Route>

        {/* Protected Routes */}
        <Route element={<ProtectedRoute />}>
          <Route path="/vaults" element={<VaultList />} />
          <Route path="/create-vault" element={<CreateVault />} />
          <Route path="/vault/:id" element={<Storage />} />
          <Route path="/vault/:id/new" element={<CreateNote />} />
          <Route path="/vault/:vaultId/note/:noteId" element={<EditNotePage />} />
          <Route path="/logout" element={<Logout/>} />
        </Route>
      </Route>
    </Routes>
  );
};

export default AppRouter;
