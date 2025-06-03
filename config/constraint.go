package config

var PromptDirectives = `
You are an AI expert specializing in generating internet policies, including Cookie Policies, Terms and Conditions, and Privacy Policies. Based on the provided user data and the specified policy type, generate a comprehensive, professional policy document in MARKDOWN format.

INSTRUCTIONS:
1. **Determine the Policy Type**: The user will specify whether the document is a **Cookie Policy**, **Terms and Conditions**, or **Privacy Policy**. Generate the appropriate document based on this input.
2. **Use Professional Legal Language**: Ensure the document is written in clear, professional legal language suitable for websites and applications.
3. **Include Relevant Sections**: Structure the document with the following sections, including only those relevant to the specified policy type and provided user data:
   - **Introduction**: Briefly explain the purpose of the policy.
   - **Definitions**: Clarify key terms used in the policy (e.g., "cookies," "personal data," "user").
   - **User Rights and Responsibilities**: For Terms and Conditions, outline user rights and obligations.
   - **Data Collection and Use**: For Privacy Policies, detail how user data is collected, used, and protected.
   - **Cookie Usage**: For Cookie Policies, describe the types of cookies used, their purposes, and how users can manage them.
   - **Liability and Disclaimers**: For Terms and Conditions, include limitations of liability and disclaimers.
   - **Third-Party Services**: If applicable, mention any third-party services (e.g., analytics, advertising) relevant to the policy.
   - **Your Rights and Choices**: Explain users' rights (e.g., data access, cookie consent) and how to exercise them.
   - **Contact Information**: Include contact details if provided in the user data.
   - **Updates to This Policy**: Explain how and when the policy may be updated.
4. **Structure with Markdown Headings**: Use proper Markdown headings (#, ##, ###) to organize the document.
5. **Ensure Compliance**: If the user data mentions specific regulations (e.g., GDPR, CCPA), ensure the policy addresses compliance with those regulations.
6. **Tailor to User Data**: Include only the categories and details specified in the user data. Omit sections for which no data is provided.
7. **Format as Clean Markdown**: Ensure the document is well-formatted and easy to read.
8. **Effective Date and Last Updated**: Include these only if provided in the user data.

REQUIREMENTS:
- Use professional, clear, and legally appropriate language.
- Include proper legal disclaimers where necessary.
- Use a table format for detailed lists (e.g., cookie details, data types).
- Ensure the document is ready for implementation on websites and applications.

USER DATA:
%s

Based on the provided user data and the specified policy type, generate the complete policy document in Markdown format. Ensure it is comprehensive, legally appropriate, and tailored to the user's needs.
`

var DummyResult = `
# Cookie Policy for Hedda Franks

## Introduction

This Cookie Policy ("Policy") explains how Clark Conley LLC ("we," "us," or "our") uses cookies and similar tracking technologies on the website [https://www.cuwybuwy.me.uk](https://www.cuwybuwy.me.uk) (the "Site"). By accessing or using the Site, you consent to the use of cookies as described in this Policy. 

This Policy is designed to comply with the **General Data Protection Regulation (GDPR)** and other applicable laws in **Aspernatur vel velit**. 

---

## What Are Cookies?

Cookies are small text files stored on your device when you visit a website. They enable the Site to remember your actions and preferences (e.g., login details, language preferences) over time, improving your browsing experience.

---

## Types of Cookies We Use

Based on your interaction with our Site, we use the following categories of cookies:

| **Category**       | **Purpose**                                                                 | **Duration** |
|--------------------|-----------------------------------------------------------------------------|-------------|
| Necessary Cookies | Essential for the Site to function (e.g., security, load balancing).        | Session     |

**Notes:**
- **Necessary Cookies** are always active and do not require user consent under GDPR.
- We do not use preference, analytics, advertising, social media, content, or personalization cookies at this time.

---

## How We Use Cookies

Cookies on our Site are used for the following purposes:
- **Site Functionality**: To ensure the Site operates correctly and securely.
- **First-Party Origin**: All cookies are set by us (no third-party cookies are used).

---

## Cookie Management and Consent

### Consent Mechanism
We respect your privacy rights and provide the following options for managing cookies:
- **Category Selection**: You may adjust cookie preferences via our cookie consent tool (if enabled).
- **Browser Settings**: You can disable or delete cookies through your browser settings. However, blocking necessary cookies may affect Site functionality.

**Instructions for Browser Settings:**
- Chrome: Settings > Privacy and Security > Cookies and Site Data
- Firefox: Options > Privacy & Security > Cookies and Site Data
- Safari: Preferences > Privacy > Manage Website Data

---

## Your Rights and Choices

Under GDPR, you have the following rights regarding cookies:
1. **Right to Access**: Request details about the cookies we use.
2. **Right to Withdraw Consent**: Adjust preferences or disable non-essential cookies.
3. **Right to Complain**: Lodge a complaint with a supervisory authority in **Aspernatur vel velit**.

---

## Contact Information

For questions or requests related to this Policy, contact us via:
- **Email**: [hynomicisu@mailinator.com](mailto:hynomicisu@mailinator.com)
- **Phone**: +1 (342) 352-7733
- **Mailing Address**: Consequuntur ducimus

---

## Updates to This Policy

We may update this Policy periodically to reflect changes in our practices or legal requirements. The latest version will always be posted on this page.
`
