import { fetchTodosPending, fetchTodosSuccess, fetchTodosFailure } from "./todosSlice";
import { AppDispatch } from "../store";
import Cookies from "js-cookie";

export const fetchTodos = () => async (dispatch: AppDispatch) => {
    dispatch(fetchTodosPending());
    const accessToken = Cookies.get("accessToken");
    const refreshToken = Cookies.get("refreshToken");
    try {
        const response = await fetch("http://localhost:3001/api/v1/todos?filter=", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                ...(accessToken && { access_token: accessToken }),
                ...(refreshToken && { refresh_token: refreshToken }),
            },
        });
        if (!response.ok) throw new Error("Failed to fetch todos");
        const data = await response.json();
        dispatch(fetchTodosSuccess(data));
    } catch (error) {
        dispatch(fetchTodosFailure());
    }
};
