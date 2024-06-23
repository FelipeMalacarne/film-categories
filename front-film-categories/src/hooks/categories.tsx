import axios from "axios";
import { BaseURL } from "../config";
import { Category } from "../types";
import { useEffect, useState } from "react";
import { useToast } from "../components/ui/use-toast";

export const useCategories = () => {
  const { toast } = useToast();
  const [categories, setCategories] = useState<Category[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    getCategories();
  }, []);

  const getCategories = async () => {
    setIsLoading(true);
    const response = await axios.get<Category[]>(BaseURL + "/category");
    const data = response.data;

    setCategories(data);
    setIsLoading(false);
  };

  const createCategory = async (category: Partial<Category>) => {
    try {
      setIsLoading(true);
      const response = await axios.post<Category>(BaseURL + "/category", category);
      toast({
        title: "Category created succesfully",
        description: "The category was created successfully",
      });
      setCategories([...categories, response.data]);
    } catch (error: any) {
      toast({
        title: "Error creating category",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  const deleteCategory = async (id: string) => {
    try {
      setIsLoading(true);
      await axios.delete(BaseURL + "/category/" + id);
      setCategories(categories.filter((category) => category.id !== id));
      toast({
        title: "Category deleted succesfully",
        description: "The category was deleted successfully",
      });
    } catch (error: any) {
      toast({
        title: "Error deleting category",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  const updateCategory = async (id: string, category: Partial<Category>) => {
    try {
      setIsLoading(true);
      const response = await axios.put<Category>(BaseURL + "/category/" + id, category);
      setCategories(categories.map((category) => category.id === id ? response.data : category));
      toast({
        title: "Category updated succesfully",
        description: "The category was updated successfully",
      });
    } catch (error: any) {
      toast({
        title: "Error updating category",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  return {
    categories,
    isLoading,
    createCategory,
    deleteCategory,
    updateCategory,
  };
}