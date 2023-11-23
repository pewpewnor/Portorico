import TestTemplate from "./TestTemplate";

export const templateNames = ["Test Template", "Test Template 2"];

export function getTemplate(templateName: string, content: any) {
	if (templateName === "Test Template") {
		return <TestTemplate {...content} isEditing={false} />;
	}
	return undefined;
}

export function getTemplateForEditing(templateName: string, content: any) {
	if (templateName === "Test Template") {
		return <TestTemplate {...content} isEditing={true} />;
	}
	return undefined;
}
