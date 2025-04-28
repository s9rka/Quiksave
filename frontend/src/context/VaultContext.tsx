import { useParams } from "react-router-dom";

export const useVault = () => {
  const { id, vaultId } = useParams<{ id: string; vaultId: string }>();
  // Use vaultId from nested route if available, otherwise use id from top-level route
  const vaultIdNum = vaultId ? Number(vaultId) : (id ? Number(id) : null);

  return {
    vaultId: vaultIdNum,
  };
};
