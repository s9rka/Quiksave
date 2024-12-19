import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { Label } from "../ui/label";
import { useAuthServices } from "@/hooks/useAuthServices";
import { useNavigate } from "react-router-dom";

export default function LoginForm() {
  const { loginMutation } = useAuthServices();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    loginMutation.mutate(
      { username, password },
      {
        onSuccess: (data) => {
          console.log("Access Token:", data.accessToken);
          localStorage.setItem("accessToken", data.accessToken);
          navigate(`/${username}`);
        },
        onError: (error) => {
          console.error("Login failed:", error.message);
        },
      }
    );
  };

  return (
    <div className="max-w-sm mx-auto mt-16 p-6 border rounded-lg shadow-sm text-left">
      <h2 className="text-2xl font-bold mb-4">Login</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        {loginMutation.isError && (
          <p className="text-red-500 text-sm">
            {(loginMutation.error as Error).message}
          </p>
        )}
        <div>
          <Label htmlFor="username" className="block text-sm font-medium mb-2">
            Username
          </Label>
          <Input
            id="username"
            type="username"
            placeholder="Enter your username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div>
          <Label htmlFor="password" className="block text-sm font-medium mb-2">
            Password
          </Label>
          <Input
            id="password"
            type="password"
            placeholder="Enter your password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <Button
          type="submit"
          className="w-full"
          disabled={loginMutation.isPending}
        >
          {loginMutation.isPending ? "Logging in..." : "Login"}
        </Button>
      </form>
    </div>
  );
}
