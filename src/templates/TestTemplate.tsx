interface TestTemplateProps {
	editing: boolean;
	text1: string;
	text2: string;
}

export default function TestTemplate(props: TestTemplateProps) {
	return (
		<div>
			<h1>{props.text1}</h1>
			<hr />
			<p>{props.text2}</p>
		</div>
	);
}
