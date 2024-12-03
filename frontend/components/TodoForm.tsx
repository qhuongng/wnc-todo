"use client";

import { useForm, SubmitHandler } from "react-hook-form";
import { Icon } from "@iconify/react";

import { addTodo } from "@/lib/redux/todosSlice";
import Cookies from "js-cookie";
import { useAppDispatch } from "@/lib/hooks";
import { Todo } from "@/lib/todos";

interface TodoInput {
    content: string;
}

const TodoForm: React.FC = () => {
    const dispatch = useAppDispatch();


    const {
        register,
        handleSubmit,
        getValues,
        reset,
        formState: { errors },
    } = useForm<TodoInput>();

    const onSubmit: SubmitHandler<TodoInput> = async (data) => {
        try {
            const accessToken = Cookies.get("accessToken");
            const refreshToken = Cookies.get("refreshToken");
            const dataToAdd = {
                content: data.content,
                completed: false
            }
            const response = await fetch("http://localhost:3001/api/v1/todos", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    ...(accessToken && { access_token: accessToken }),
                    ...(refreshToken && { refresh_token: refreshToken }),
                },
                body: JSON.stringify(dataToAdd),
            });

            if (response.ok) {
                reset();
                const responseData = await response.json();
                const newTodo: Todo = {
                    id: responseData.data.id,
                    content: responseData.data.content,
                    completed: responseData.data.completed
                }
                dispatch(addTodo(newTodo));
            } else {
                const data = await response.json();
                alert(data.message || "Đã xảy ra lỗi trong quá trình thêm. Hãy thử lại sau");
            }
        } catch (error) {
            alert("Đã xảy ra lỗi trong quá trình thêm. Hãy thử lại sau");
        }

    };

    return (
        <div>
            <label htmlFor="add-todo" className="btn btn-secondary btn-square ml-8">
                <Icon icon="ic:round-plus" style={{ fontSize: "28px" }} />
            </label>

            <input type="checkbox" id="add-todo" className="modal-toggle" />

            <div className="modal" role="dialog">
                <div className="modal-box">
                    <h2 className="font-bold" style={{ marginTop: 0 }}>
                        Create a new to-do
                    </h2>
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <textarea
                            className="textarea textarea-bordered w-full"
                            placeholder="Buy milk"
                            {...register("content", { required: true })}
                        />
                        {errors.content && (
                            <span className="text-red-800">Type something first!</span>
                        )}

                        <div className="modal-action">
                            <label htmlFor="add-todo" className="btn">
                                Cancel
                            </label>

                            <button
                                className="btn btn-secondary"
                                type="submit"
                                onClick={() => {
                                    const contentValue = getValues("content");
                                    const addTodoCheckbox = document.getElementById("add-todo");

                                    if (contentValue.length > 0 && addTodoCheckbox) {
                                        (addTodoCheckbox as HTMLInputElement).checked = false;
                                    }
                                }}
                            >
                                Create
                            </button>
                        </div>
                    </form>
                </div>

                <label className="modal-backdrop" htmlFor="add-todo" />
            </div>
        </div>
    );
};

export default TodoForm;
