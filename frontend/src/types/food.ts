export interface FoodItem {
  id: string;
  name: string;
  description: string;
  price: number;
  image: string;
  category: string;
  salesCount?: number;
}

export interface CartItem extends FoodItem {
  quantity: number;
}
