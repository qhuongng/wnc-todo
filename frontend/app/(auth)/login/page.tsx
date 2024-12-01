"use client";

import { useRouter } from "next/navigation";
import { useForm, SubmitHandler } from "react-hook-form";
import Link from "next/link";
import { useAppDispatch } from "@/lib/hooks";
import { setUsername } from "@/lib/redux/userSlice";

interface LoginInputs {
    username: string,
    password: string
}

const Login: React.FC = () => {

    const dispatch = useAppDispatch();
    const router = useRouter();

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors },
    } = useForm<LoginInputs>();

    const onSubmit: SubmitHandler<LoginInputs> = async (data) => {
        try {
            const response = await fetch("http://localhost:3001/api/v1/users/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                const data = await response.json();
                dispatch(setUsername(data.data));
                router.push('/');
            } else {
                console.log(("Thông tin đăng nhập không hợp lệ."));
            }
        } catch (error) {
            console.log(error);
        }
        reset();
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="w-full max-w-md bg-white rounded-lg shadow-md p-6">
                <h2 className="text-2xl font-bold text-center text-gray-800">Đăng nhập</h2>
                <form className="mt-4" onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-4">
                        <label htmlFor="username" className="block text-sm font-medium text-gray-600">
                            Email
                            <span className="text-red-500 ml-1">*</span>
                        </label>
                        <input
                            type="text"
                            id="username"
                            className="w-full mt-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Nhập tên đăng nhập"
                            {...register("username", { required: true })}
                        />
                        {errors.username && (
                            <span className="text-red-500">Tên đăng nhập không được bỏ trống!</span>
                        )}
                    </div>
                    <div className="mb-6">
                        <label htmlFor="password" className="block text-sm font-medium text-gray-600">
                            Mật khẩu
                            <span className="text-red-500 ml-1">*</span>

                        </label>
                        <input
                            type="password"
                            id="password"
                            className="w-full mt-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Nhập mật khẩu"
                            {...register("password", { required: true })}
                        />
                        {errors.password && (
                            <span className="text-red-500">Mật khẩu không được bỏ trống!</span>
                        )}
                    </div>
                    <button
                        type="submit"
                        className="w-full px-4 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400"
                    >
                        Đăng nhập
                    </button>
                </form>
                <p className="mt-4 text-sm text-center text-gray-600">
                    Chưa có tài khoản? <Link href="#" className="text-blue-500 hover:underline">Đăng ký</Link>
                </p>
            </div>
        </div>

    );
}
export default Login;