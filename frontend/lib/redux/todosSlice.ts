import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { Todo } from "@/lib/todos";

export interface TodosState {
    todos: Array<Todo>;
    filteredTodos: Array<Todo>;
    status: "idle" | "pending" | "success" | "failure";
}

const initialState: TodosState = {
    todos: [],
    filteredTodos: [],
    status: "idle",
};

export const todosSlice = createSlice({
    name: "todos",
    initialState,
    reducers: {
        setStatusPending: (state) => {
            state.status = "pending";
        },
        fetchTodosSuccess: (state, action: PayloadAction<Array<Todo>>) => {
            state.status = "success";
            state.todos = action.payload;
            state.filteredTodos = action.payload;
        },
        setStatusFailure: (state) => {
            state.status = "failure";
        },
        addTodo: (state, action: PayloadAction<string>) => {
            state.status = "success";
            const newTodo: Todo = {
                id: state.todos.length + 1,
                content: action.payload,
                completed: false,
            };

            state.todos.push(newTodo);
            state.filteredTodos.push(newTodo);
        },
        markTodo: (state, action: PayloadAction<number>) => {
            state.status = "success";
            const todo = state.todos.find((todo: Todo) => todo.id === action.payload);
            const todoFromFiltered = state.filteredTodos.find(
                (todo: Todo) => todo.id === action.payload
            );

            if (todo) {
                todo.completed = !todo.completed;
                todoFromFiltered!.completed = !todoFromFiltered!.completed;
            }
        },
        filterTodo: (state, action: PayloadAction<string>) => {
            state.status = "success";
            state.filteredTodos = state.todos.filter((todo: Todo) => {
                return todo.content.toLowerCase().includes(action.payload.toLowerCase());
            });
        },
    },
});

export const {
    fetchTodosSuccess,
    setStatusPending,
    setStatusFailure,
    addTodo,
    markTodo,
    filterTodo,
} = todosSlice.actions;
export const todosReducer = todosSlice.reducer;
