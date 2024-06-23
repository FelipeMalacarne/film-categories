import axios from "axios";
import { BaseURL } from "../config";
import { Film } from "../types";
import { useEffect, useState } from "react";
import { useToast } from "../components/ui/use-toast";

export const useFilms = () => {
  const { toast } = useToast();
  const [films, setFilms] = useState<Film[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    getFilms();
  }, []);

  const getFilms = async () => {
    setIsLoading(true);
    const response = await axios.get<Film[]>(BaseURL + "/film");
    const data = response.data;

    setFilms(data);
    setIsLoading(false);
  };

  const createFilm = async (film: Partial<Film>) => {
    try {
      setIsLoading(true);
      const response = await axios.post<Film>(BaseURL + "/film", film);
      toast({
        title: "Film created succesfully",
        description: "The film was created successfully",
      });
      setFilms([...films, response.data]);
    } catch (error: any) {
      toast({
        title: "Error creating film",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  const deleteFilm = async (id: string) => {
    try {
      setIsLoading(true);
      await axios.delete(BaseURL + "/film/" + id);
      setFilms(films.filter((film) => film.id !== id));
      toast({
        title: "Film deleted succesfully",
        description: "The film was deleted successfully",
      });
    } catch (error: any) {
      toast({
        title: "Error deleting film",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  const updateFilm = async (id: string, film: Partial<Film>) => {
    try {
      setIsLoading(true);
      const response = await axios.put<Film>(BaseURL + "/film/" + id, film);
      setFilms(films.map((f) => (f.id === id ? response.data : f)));
      toast({
        title: "Film updated succesfully",
        description: "The film was updated successfully",
      });
    } catch (error: any) {
      toast({
        title: "Error updating film",
        description: error.response.data.message,
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  return { films, isLoading, createFilm, deleteFilm, updateFilm }
};
