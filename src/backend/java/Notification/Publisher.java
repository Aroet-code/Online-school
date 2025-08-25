package Notification;

import User.User;

import java.util.ArrayList;

public class Publisher {
    public void update(ArrayList<User> users){
        for (User user : users){
            user.onNotification();
        }
    }
    public void sendNotification(String notification, ArrayList<User> users,
                                 ArrayList<String> tags, boolean mustHaveAllTheTags){
        for (User user : users){
            if (mustHaveAllTheTags){
                for (String tag : tags){
                    if (!user.hasTag(tag)){
                        break;
                    } else {
                        NotificationService.update(user);
                    }
                }
            } else {
                for (String tag : tags){
                    if (user.hasTag(tag)){
                        NotificationService.update(user);
                        break;
                    }
                }
            }
        }
    }
}
