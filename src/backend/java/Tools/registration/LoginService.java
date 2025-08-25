package Tools.registration;

import Notification.EmailNotificationService;
import Notification.NotificationService;
import User.User;

import java.util.ArrayList;

public class LoginService{
    public static void login(String username, String password, ArrayList<LoginUser> users)
            throws LoginErrorException{
        boolean found = false;
        try {
            for (LoginUser user : users) {
                    if (user.getUsername().equals(username)) {
                        // username found
                        found = true;
                        if (user.getPassword().equals(password)) {
                            // the password is correct
                            break;
                        } else {
                            throw new LoginErrorException("Password incorrect!");
                        }
                    }
                }
            if (!found){
                throw new LoginErrorException("No user with such username found!");
            }
            }   catch (LoginErrorException e) {
            System.out.println(e.getMessage());
        }
    }

    public static void register(ArrayList<User> users, ArrayList<LoginUser> loginUsers,
                                String username, String password, String email, String phoneNumber)
    throws RegisterErrorException{
        String userName = username.toLowerCase();
        try {
            for (LoginUser user : loginUsers){
                if (user.getUsername().equals(userName)){
                    throw new RegisterErrorException("Username already taken!");
                }
            }
            users.add(new User(userName));
            loginUsers.add(new LoginUser(userName, password, email, phoneNumber));
        } catch (RegisterErrorException e) {
            System.out.println(e.getMessage());
        }
    }

    public static void emailRecovery(ArrayList<LoginUser> loginUsers, String email)
            throws LoginErrorException{
        try {
            boolean found = false;
            for (LoginUser user : loginUsers) {
                if (user.getEmail().equals(email)){
                    EmailNotificationService emailNotificationService = new EmailNotificationService();
                    emailNotificationService.onUpdate(user, "Blah, blah, blah");
                    // send recovery notification here
                    found = true;
                    break;
                }
            }
            if (!found){
                throw new LoginErrorException("No user with that email found!");
            }
        } catch (LoginErrorException e) {
            System.out.println(e.getMessage());
        }
    }

    public static void usernameRecovery(ArrayList<LoginUser> loginUsers, String username)
            throws LoginErrorException{
        try {
            boolean found = false;
            for (LoginUser user : loginUsers){
                if (user.getUsername().equals(username)){
                    // send recovery notification here
                    found = true;
                    break;
                }
            }
            if (!found){
                throw new LoginErrorException("No user with such username found!");
            }
        } catch (LoginErrorException e){
            System.out.println(e.getMessage());
        }
    }

    public static void phoneNumberRecovery(ArrayList<LoginUser> loginUsers, String phoneNumber)
            throws LoginErrorException{
        try {
            boolean found = false;
            for (LoginUser user : loginUsers){
                if (user.getPhoneNumber().equals(phoneNumber)){
                    // send recovery notification here
                    found = true;
                    break;
                }
            }
            if (!found){
                throw new LoginErrorException("No user with that phone number found!");
            }
        } catch (LoginErrorException e) {
            System.out.println(e.getMessage());
        }
    }
}
