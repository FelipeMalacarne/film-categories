import { useEffect, useState } from "react";
import { Film } from "../types";
import { BaseURL } from "../config";
import axios from "axios";
import { useToast } from "../components/ui/use-toast";

export const useFilms = () => {
  const { toast } = useToast();
  const [films, setFilms] = useState<Film[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const getFilms = async () => {
    setIsLoading(true);
    const response = await axios.get<Film[]>(BaseURL + "/film");
    const data = response.data;

    setFilms(data);
    setIsLoading(false);
  };

  useEffect(() => {
    getFilms();
  }, []);

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

  return { films, createFilm, isLoading };
};
