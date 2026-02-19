import axios from "axios";


export const getComboRecommendations = async (itemId: string) => {
const res = await axios.get(`http://localhost:5000/api/v1/food/recommend?itemId=${itemId}`);
return res.data.data || [];
};


export const searchFoodItems = async (query: string) => {
const res = await axios.get(`http://localhost:5000/api/v1/search?query=${encodeURIComponent(query)}`);
return res.data.results || [];
};