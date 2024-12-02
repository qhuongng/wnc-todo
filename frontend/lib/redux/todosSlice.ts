import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { Todo } from "@/lib/todos";

export interface TodosState {
    todos: Array<Todo>;
    filteredTodos: Array<Todo>;
    status: string,
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
        fetchTodosPending: (state) => {
            state.status = 'pending';
        },
        fetchTodosSuccess: (state, action: PayloadAction<Array<Todo>>) => {
            state.todos = action.payload;
            state.filteredTodos = action.payload;
        },
        fetchTodosFailure: (state) => {
            state.status = 'idle';
        },
        addTodo: (state, action: PayloadAction<string>) => {
            const newTodo: Todo = {
                id: state.todos.length + 1,
                content: action.payload,
                completed: false,
            };

            state.todos.push(newTodo);
            state.filteredTodos.push(newTodo);
        },
        markTodo: (state, action: PayloadAction<number>) => {
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
            state.filteredTodos = state.todos.filter((todo: Todo) => {
                return todo.content.toLowerCase().includes(action.payload.toLowerCase());
            });
        },
    },
});

export const { fetchTodosSuccess, fetchTodosFailure, fetchTodosPending, addTodo, markTodo, filterTodo } = todosSlice.actions;
export const todosReducer = todosSlice.reducer;
