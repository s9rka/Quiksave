import LoginForm from "@/components/auth/Login";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <>
      <LoginForm />
    </>
  );
}
