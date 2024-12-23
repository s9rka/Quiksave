import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import "./App.css";
import AppRouter from "./routes/Router";
import NavDrawer from "./components/layout/NavDrawer";
import { isAuthenticatedAtom } from "./services/auth";
import { useAtom } from "jotai";

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
  const [isAuthenticated] = useAtom(isAuthenticatedAtom);

  return (
    <QueryClientProvider client={queryClient}>
      <AppRouter />
      {isAuthenticated && <NavDrawer />}
    </QueryClientProvider>
  );
}

export default App;
