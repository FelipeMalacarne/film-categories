export interface Film {
  id: string;
  name: string;
  author: string;
  description: string;
  duration: number;
  release_date: Date;
  created_at: Date;
  updated_at: Date;
}

export interface Category {
  id: string;
  name: string;
  created_at: Date;
  updated_at: Date;
}

export interface Supplier {
  id: string;
  name: string;
  email: string;
  phone: string;
  created_at: Date;
  updated_at: Date;
}
