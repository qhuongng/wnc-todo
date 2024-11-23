import TodoItem from "./TodoItem";

import { useAppSelector } from "@/lib/hooks";
import { RootState } from "@/lib/store";

const TodoList = () => {
    const todos = useAppSelector((state: RootState) => state.todos.filteredTodos);

    return (
        <div className="flex flex-col pb-8 px-8">
            {todos.map((todo) => (
                <TodoItem key={todo.id} todo={todo} />
            ))}
        </div>
    );
};

export default TodoList;
