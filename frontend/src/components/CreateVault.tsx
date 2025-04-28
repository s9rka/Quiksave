import { useState } from 'react';
import { VaultFormData } from '@/lib/types';
import { useCreateVault } from '@/services/mutations';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Textarea } from './ui/textarea';
import { Button } from './ui/button';

const CreateVault = () => {
    const [formData, setFormData] = useState<VaultFormData>({
        name: '',
        description: ''
    });
    const [error, setError] = useState<string>('');

    const createVaultMutation = useCreateVault();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');

        try {
            await createVaultMutation.mutateAsync(formData);
        } catch (err) {
            setError('Failed to create vault. Please try again.');
            console.error('Error creating vault:', err);
        }
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    return (
        <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl font-bold mb-6 text-center">Create New Vault</h2>
            {error && (
                <div className="mb-4 p-3 bg-red-100 text-red-700 rounded">
                    {error}
                </div>
            )}
            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <Label htmlFor="name" className="block text-sm font-medium text-gray-700">
                        Vault Name
                    </Label>
                    <Input
                        type="text"
                        id="name"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        required
                        className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                    />
                </div>
                <div>
                    <Label htmlFor="description" className="block text-sm font-medium text-gray-700">
                        Description
                    </Label>
                    <Textarea
                        id="description"
                        name="description"
                        value={formData.description}
                        onChange={handleChange}
                        required
                        rows={4}
                        className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                    />
                </div>
                <Button
                    type="submit"
                    disabled={createVaultMutation.isPending}
                    className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:bg-indigo-300"
                >
                    {createVaultMutation.isPending ? 'Creating...' : 'Create Vault'}
                </Button>
            </form>
        </div>
    );
};

export default CreateVault; 