package Notification;

import Notification.waysOfNotification.EmailNotification;
import Tools.registration.LoginUser;
import User.User;

public class EmailNotificationService implements Subscription {


    private void sendNotification(User user, EmailNotification notification){

    }

    @Override
    public void onUpdate(LoginUser user, String notificationID) {
        boolean sent = false;
        //
        // the email letter should be sent there
        //
        if (!sent){
            throw new NotificationErrorException("No email notification sent!");
        }
    }
}
