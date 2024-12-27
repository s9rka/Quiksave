import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import "./App.css";
import AppRouter from "./routes/Router";
import { useInitializeAuth } from "./context/UserContext";

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5,

      refetchOnMount: false,
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  
  const loading = useInitializeAuth();

  // Show a spinner or loading state until done
  if (loading) {
    return <div>Loading...</div>;
  }
  return (
    <QueryClientProvider client={queryClient}>
        <AppRouter />
    </QueryClientProvider>
  );
}

export default App;
