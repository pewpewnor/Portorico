import { User } from "@/types/model";
import { Dispatch, SetStateAction, createContext } from "react";

const AuthContext = createContext<
	[User | null, Dispatch<SetStateAction<User | null>>]
>([null, () => {}]);

export default AuthContext;
