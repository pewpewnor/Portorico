interface ErrorMessageProps {
	message: string;
}

export default function ErrorMessage(props: ErrorMessageProps) {
	return <p className="text-sm text-red-500">{props.message}</p>;
}
