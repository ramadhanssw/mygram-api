# mygram-api
An application called MyGram as an application that can save photos or make comments on other people's photos. In simple terms, this app is a social media platform for sharing photo moments.

As part of the Scalable Web Services with Golang Class Final Project, this project has various criteria that include:
1. Using library/package 
● github.com/dgrijalva/jwt-go
● golang.org/x/crypto
2. Implement certain tables and fields and these will be displayed in the database schema in the application.

## Database Scheme
![image](https://user-images.githubusercontent.com/81277898/234207621-fc82af18-c48b-4d09-8fb6-8cc06ca602b4.png)

### Database Scheme - Endpoint
1. User:
- Register [POST]
- Login [POST]

2. Photo: 
- GetAll [GET] 
- GetOne [GET] 
- CreatePhoto [POST] 
- UpdatePhoto [PUT] 
- DeletePhoto [DELETE]

3. Comment: 
- GetAll [GET] 
- GetOne [GET] 
- CreateComment [POST] 
- UpdateComment [PUT] 
- DeleteComment [DELETE]

4. Social Media: 
- GetAll [GET] 
- GetOne [GET] 
- CreateSocialMedia [POST] 
- UpdateSocialMedia [PUT] 
- DeleteSocialMedia [DELETE]

### Database Scheme - Validation
A. Validation for User table:
1. Email field
 - Validation to check for a valid email format
 - Validation so that it can be a unique index
 - Validation so that the email field cannot be empty or must be filled in
2. Username field
 - Validation so that it can be a unique index
 - Validation so that the username field cannot be empty or must be filled in
3. Password field
- Validate that the password field cannot be empty
- Validate that the password field has a minimum length of 6 characters
4. Age field
- Validation that the age field cannot be empty or must be filled in
- Validate that the age field must have at least a value above 8

B. Validation for the photo table:
1. Field title
- Validate that the title field cannot be empty
2. Field photo_url
- Validation that the photo_url field cannot be empty or must be filled in

C. Validation for social media table:
1. Field name
- Validate that the field name cannot be empty
2. Field social_media_url
- Validate that the social_media_url field cannot be empty or must be filled in.

D. Validation for comment:
1. Message field
- Validate that the message field cannot be empty
