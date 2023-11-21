"use client";

import ErrorMessage from "@/components/ui/ErrorMessage";
import LoggedInContext from "@/contexts/LoggedInContext";
import client from "@/lib/axios";
import { User } from "@/types/model";
import { Validations } from "@/types/responses";
import { Button, Checkbox, Input } from "@nextui-org/react";
import Cookies from "js-cookie";
import moment from "moment";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { ChangeEvent, useContext, useState } from "react";

type LoginResponse = {
	user?: User;
	validations?: Validations;
};

export default function LoginPage() {
	const router = useRouter();
	const [_, refreshLoggedIn] = useContext(LoggedInContext);

	const [isLoading, setIsLoading] = useState(false);
	const [validations, setValidations] = useState<Validations>({});
	const [inputData, setInputData] = useState({
		username: Cookies.get("username") ?? "",
		password: Cookies.get("password") ?? "",
		rememberMe: false,
	});

	function handleChange(event: ChangeEvent<HTMLInputElement>) {
		setInputData((prev) => ({
			...prev,
			[event.target.name]: event.target.value,
		}));
	}

	async function handleLogin() {
		setIsLoading(true);
		setValidations({});

		try {
			const res = await client.post("/login", inputData);
			const data = res.data as LoginResponse;

			if (res.status == 400 && data.validations) {
				setValidations(data.validations);
			} else if (res.status == 200 && data.user) {
				if (inputData.rememberMe) {
					const cookieConfig = {
						expires: moment(Date.now()).add(6, "hours").toDate(),
					};
					Cookies.set("username", inputData.username, cookieConfig);
					Cookies.set("password", inputData.password, cookieConfig);
				}
				refreshLoggedIn();
				router.push("/dashboard");
			}
		} catch (error) {
			setValidations({
				general: "server have troubles processing your request",
			});
		}

		setIsLoading(false);
	}

	return (
		<div className="flex h-[92vh] flex-col overflow-hidden md:flex-row">
			<Image
				src="/images/register.webp"
				placeholder="blur"
				blurDataURL="/images/register-blur.webp"
				className="h-32 backdrop-blur md:h-full md:w-1/2"
				width={800}
				height={800}
				alt="register"
			/>
			<div className="flex h-full animate-appearance-in flex-col items-center justify-center gap-10 bg-gradient-to-br from-white to-blue-200 py-16 md:w-1/2">
				<h1 className="text-5xl font-medium">Login</h1>
				<div className="mb-4 flex w-full max-w-xs flex-col gap-y-2 rounded bg-white px-8 pb-8 pt-6 shadow-md">
					<Input
						label="username"
						type="text"
						size="lg"
						className="select-none"
						name="username"
						value={inputData.username}
						onChange={handleChange}
					/>
					{validations.username && (
						<ErrorMessage message={validations.username} />
					)}
					<Input
						label="password"
						type="password"
						size="lg"
						className="mt-4 select-none"
						name="password"
						value={inputData.password}
						onChange={handleChange}
					/>
					{validations.password && (
						<ErrorMessage message={validations.password} />
					)}
					<div className="mt-4 flex">
						<Checkbox
							checked={inputData.rememberMe}
							onChange={() =>
								setInputData((prev) => ({
									...prev,
									rememberMe: !prev.rememberMe,
								}))
							}
						/>
						<p className="text-sm">Remember Me</p>
					</div>
					<Button
						disabled={isLoading}
						onPress={handleLogin}
						className="mt-4 w-full rounded border bg-indigo-600 px-8 py-4 text-center text-lg font-medium text-white hover:bg-indigo-700"
					>
						Next
					</Button>
					{validations.general && (
						<ErrorMessage message={validations.general} />
					)}
					<p className="mt-4 text-center text-sm">
						Haven&apos;t created an account?
						<Link href="/login">
							{" "}
							<span className="text-blue-600 underline hover:text-blue-300">
								Sign Up
							</span>
						</Link>
					</p>
				</div>
			</div>
		</div>
	);
}
