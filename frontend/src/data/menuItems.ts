import { FoodItem } from '@/types/food';
import pastaImg from '@/assets/food-pasta.jpg';
import steakImg from '@/assets/food-steak.jpg';
import sushiImg from '@/assets/food-sushi.jpg';
import salmonImg from '@/assets/food-salmon.jpg';
import dessertImg from '@/assets/food-dessert.jpg';
import burgerImg from '@/assets/food-burger.jpg';

export const menuItems: FoodItem[] = [
  {
    id: '1',
    name: 'Truffle Pasta',
    description: 'Homemade pasta with black truffle, parmesan, and fresh herbs',
    price: 32.99,
    image: pastaImg,
    category: 'Main Course',
    salesCount: 245,
  },
  {
    id: '2',
    name: 'Wagyu Steak',
    description: 'Premium A5 Wagyu beef with seasonal vegetables and red wine reduction',
    price: 89.99,
    image: steakImg,
    category: 'Main Course',
    salesCount: 198,
  },
  {
    id: '3',
    name: 'Omakase Sushi',
    description: "Chef's selection of premium sushi and sashimi",
    price: 78.99,
    image: sushiImg,
    category: 'Main Course',
    salesCount: 312,
  },
  {
    id: '4',
    name: 'Grilled Salmon',
    description: 'Wild-caught salmon with lemon butter sauce and asparagus',
    price: 42.99,
    image: salmonImg,
    category: 'Main Course',
    salesCount: 167,
  },
  {
    id: '5',
    name: 'Chocolate Soufflé',
    description: 'Warm chocolate soufflé with gold leaf and vanilla ice cream',
    price: 18.99,
    image: dessertImg,
    category: 'Dessert',
    salesCount: 289,
  },
  {
    id: '6',
    name: 'Gourmet Burger',
    description: 'Angus beef burger with truffle aioli, aged cheddar, and caramelized onions',
    price: 28.99,
    image: burgerImg,
    category: 'Main Course',
    salesCount: 421,
  },
];
