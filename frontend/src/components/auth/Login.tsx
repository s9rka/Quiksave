import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { Label } from "../ui/label";
import { useAuthServices } from "@/hooks/useAuthServices";

export default function LoginForm() {
  const { login, loading, error } = useAuthServices();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const token = await login({ username, password });
    if (token) {
      console.log("Access Token:", token);
      localStorage.setItem("accessToken", token);
    }
  };

  return (
    <div className="max-w-sm mx-auto mt-16 p-6 border rounded-lg shadow-sm text-left">
      <h2 className="text-2xl font-bold mb-4">Login</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        {error && <p className="text-red-500 text-sm">{error}</p>}
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
        <Button type="submit" className="w-full">
          Login
        </Button>
      </form>
    </div>
  );
}