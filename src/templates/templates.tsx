import BasicLanding from "./Basic/Landing";

export const templateNames = ["Test Template", "Test Template 2"];

export function getTemplate(templateName: string, content: any) {
	if (templateName === "Test Template") {
		return <BasicLanding {...content} isEditing={false} />;
	}
	return undefined;
}

export function getTemplateForEditing(
	templateName: string,
	content: any,
	changeField: (fieldName: string, value: string) => void
) {
	if (templateName === "Test Template") {
		return (
			<BasicLanding
				{...content}
				isEditing={true}
				changeField={changeField}
			/>
		);
	}
	return undefined;
}
