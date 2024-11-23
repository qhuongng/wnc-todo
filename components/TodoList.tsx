import TodoItem from "./TodoItem";
import { Todo } from "@/lib/todos";

interface TodoListProps {
    todos: Todo[];
    markTodo: (id: number) => void;
}

const TodoList: React.FC<TodoListProps> = ({ todos, markTodo }) => {
    return (
        <div className="flex flex-col pb-8 px-8">
            {todos.map((todo: Todo) => (
                <TodoItem key={todo.id} todo={todo} markTodo={markTodo} />
            ))}
        </div>
    );
};

export default TodoList;
