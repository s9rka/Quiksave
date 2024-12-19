import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import "./App.css";
import LoginForm from "./components/auth/Login";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div>
        <LoginForm />
      </div>
      <ReactQueryDevtools initialIsOpen={false} />

    </QueryClientProvider>
  );
}

export default App;
