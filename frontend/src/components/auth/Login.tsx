import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "../ui/label";
import { useLogin } from "@/services/mutations";
import { LoginCredentials } from "@/lib/types";
import { useForm, SubmitHandler } from "react-hook-form";
import { useNavigate } from "react-router-dom";

export default function LoginForm() {
  const loginMutation = useLogin();
  const { register, handleSubmit } = useForm<LoginCredentials>();

  const navigate = useNavigate();

  const onSubmit: SubmitHandler<LoginCredentials> = (data) => {
    loginMutation.mutate(data);
    navigate("/vaults")
  };

  return (
    <div className="max-w-sm mx-auto mt-16 p-6 border rounded-lg shadow-sm text-left">
      <h2 className="text-2xl font-bold mb-4">Login</h2>
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
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
            type="text"
            {...register("username", {required: true})}
            placeholder="Enter your username"
          />
        </div>
        <div>
          <Label htmlFor="password" className="block text-sm font-medium mb-2">
            Password
          </Label>
          <Input
            id="password"
            type="password"
            {...register("password", {required: true})}
            placeholder="Enter your password"
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
