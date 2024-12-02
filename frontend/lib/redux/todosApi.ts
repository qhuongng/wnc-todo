import { fetchTodosPending, fetchTodosSuccess, fetchTodosFailure } from "./todosSlice";
import { AppDispatch } from "../store";

export const fetchTodos = () => async (dispatch: AppDispatch) => {
  dispatch(fetchTodosPending());
  try {
    const response = await fetch('http://localhost:30001/api/v1/');
    if (!response.ok) throw new Error('Failed to fetch todos');
    const data = await response.json();
    dispatch(fetchTodosSuccess(data));
  } catch (error) {
    dispatch(fetchTodosFailure());
  }
};
