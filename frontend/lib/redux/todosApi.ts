import { setStatusPending, setStatusFailure, fetchTodosSuccess, markTodo } from "./todosSlice";
import { AppDispatch } from "../store";
import Cookies from "js-cookie";

export const fetchTodos = () => async (dispatch: AppDispatch) => {
    dispatch(setStatusPending());
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
        const responseData = await response.json();
        dispatch(fetchTodosSuccess(responseData.data));
    } catch (error) {
        dispatch(setStatusFailure());
    }
};

export const updateTodo = (todo: { id: number; content: string; completed: boolean }) => async (dispatch: AppDispatch) => {
    dispatch(setStatusPending());
    const accessToken = Cookies.get("accessToken");
    const refreshToken = Cookies.get("refreshToken");

    try {
        const response = await fetch(`http://localhost:3001/api/v1/todos/${todo.id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                ...(accessToken && { access_token: accessToken }),
                ...(refreshToken && { refresh_token: refreshToken }),
            },
            body: JSON.stringify({ ...todo, completed: !todo.completed }),
        });

        if (!response.ok) {
            dispatch(setStatusFailure());
        }

        dispatch(markTodo(todo.id)); 
    } catch (error) {
        dispatch(setStatusFailure());
    }
};
