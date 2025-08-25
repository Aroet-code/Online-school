package Notification;

import Tools.registration.LoginUser;
import User.User;

public interface Subscription {
    void onUpdate(LoginUser user, String notificationID);
}
