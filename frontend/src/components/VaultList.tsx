import { useNavigate } from 'react-router-dom';
import { useVaults } from '@/services/queries';
import { Button } from './ui/button';

const VaultList = () => {
    const navigate = useNavigate();
    const { data, isLoading, error } = useVaults();
    
    // Ensure vaults is always an array, even if data is null or undefined
    const vaults = data || [];

    const handleCreateVault = () => {
        navigate('/create-vault');
    };

    const formatDate = (dateString: string) => {
        return new Date(dateString).toLocaleDateString();
    };

    if (isLoading) {
        return (
            <div className="flex justify-center items-center h-64">
                <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
            </div>
        );
    }

    return (
        <div className="max-w-4xl mx-auto mt-10 p-6">
            <div className="flex justify-between items-center mb-6">
                <h2 className="text-2xl font-bold">Your Vaults</h2>
                {vaults.length > 0 &&
                <Button
                    onClick={handleCreateVault}
                    className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                    Create New Vault
                </Button>
                }
            </div>

            {error && (
                <div className="mb-4 p-3 bg-red-100 text-red-700 rounded">
                    Failed to fetch vaults. Please try again.
                </div>
            )}

            {!error && vaults.length === 0 ? (
                <div className="text-center py-12">
                    <p className="text-gray-500 mb-4">You don't have any vaults yet.</p>
                    <Button
                        onClick={handleCreateVault}
                        className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    >
                        Create Your First Vault
                    </Button>
                </div>
            ) : (
                !error && (
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                        {vaults.map((vault) => (
                            <div
                                key={vault.id}
                                onClick={() => navigate(`/vault/${vault.id}`)}
                                className="p-6 bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow cursor-pointer"
                            >
                                <h3 className="text-xl font-semibold mb-2">{vault.name}</h3>
                                <p className="text-gray-600 mb-4">{vault.description}</p>
                                <p className="text-sm text-gray-500">
                                    Created on {formatDate(vault.created_at)}
                                </p>
                            </div>
                        ))}
                    </div>
                )
            )}
        </div>
    );
};

export default VaultList;