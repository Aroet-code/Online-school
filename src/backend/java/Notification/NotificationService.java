package Notification;

import Tools.registration.LoginUser;

import java.util.ArrayList;

public class NotificationService {
    public static void update(LoginUser user, String notificationID) throws NotificationErrorException {
        ArrayList<Subscription> subscriptions = user.getSubscriptions();
        for (Subscription subscription : subscriptions){
            try {
                subscription.onUpdate(user, notificationID);
            } catch (NotificationErrorException e) {
                System.out.println(e.getMessage());
            }
        }
    }
}
