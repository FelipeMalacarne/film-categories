import axios from "axios";
import { BaseURL } from "../config";
import { Supplier } from "../types";
import { useEffect, useState } from "react";
import { useToast } from "../components/ui/use-toast";

export const useSuppliers = () => {
  const { toast } = useToast();
  const [suppliers, setSuppliers] = useState<Supplier[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    getSuppliers();
  }, []);

  const getSuppliers = async () => {
    setIsLoading(true);
    const response = await axios.get<Supplier[]>(BaseURL + "/supplier");
    const data = response.data;

    setSuppliers(data);
    setIsLoading(false);
  };

  const createSupplier = async (supplier: Partial<Supplier>) => {
    try {
      setIsLoading(true);
      const response = await axios.post<Supplier>(BaseURL + "/supplier", supplier);
      toast({
        title: "Supplier created succesfully",
        description: "The supplier was created successfully",
      });
      setSuppliers([...suppliers, response.data]);
    } catch (error: any) {
      toast({
        title: "Error creating supplier",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  const deleteSupplier = async (id: string) => {
    try {
      setIsLoading(true);
      await axios.delete(BaseURL + "/supplier/" + id);
      setSuppliers(suppliers.filter((supplier) => supplier.id !== id));
      toast({
        title: "Supplier deleted succesfully",
        description: "The supplier was deleted successfully",
      });
    } catch (error: any) {
      toast({
        title: "Error deleting supplier",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  const updateSupplier = async (id: string, supplier: Partial<Supplier>) => {
    try {
      setIsLoading(true);
      const response = await axios.put<Supplier>(BaseURL + "/supplier/" + id, supplier);
      setSuppliers(suppliers.map((s) => s.id === id ? response.data : s));
      toast({
        title: "Supplier updated succesfully",
        description: "The supplier was updated successfully",
      });
    } catch (error: any) {
      toast({
        title: "Error updating supplier",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  return { suppliers, isLoading, getSuppliers, createSupplier, deleteSupplier, updateSupplier };
}