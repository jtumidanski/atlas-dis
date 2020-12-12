package com.atlas.dis.processor;

import com.atlas.dis.database.administrator.MonsterDropAdministrator;
import com.atlas.dis.database.provider.MonsterDropProvider;
import com.atlas.dis.model.MonsterDropData;
import com.atlas.dis.rest.attribute.MonsterDropAttributes;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import database.Connection;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public final class MonsterDropProcessor {
   private MonsterDropProcessor() {
   }

   public static void init() {
      int dropCount = Connection.instance()
            .list(MonsterDropProvider::getAll)
            .size();
      if (dropCount == 0) {
         BufferedReader fileReader;
         try {
            fileReader = new BufferedReader(new FileReader("/service/drop_data.json"));
         } catch (FileNotFoundException e) {
            System.out.println("Unable to locate drop_data.json");
            return;
         }
         List<MonsterDropData> dropData = new ArrayList<>();
         try {
            String line = fileReader.readLine();
            while (line != null) {
               dropData.add(processLine(line));
               line = fileReader.readLine();
            }
         } catch (IOException e) {
            e.printStackTrace();
         }
         Connection.instance().with(entityManager -> MonsterDropAdministrator.createBulk(entityManager, dropData));
      }
   }

   protected static MonsterDropData processLine(String line) {
      ObjectMapper objectMapper = new ObjectMapper();
      try {
         MonsterDropAttributes attributes = objectMapper.readValue(line, MonsterDropAttributes.class);
         return new MonsterDropData(0, attributes.monsterId(), attributes.itemId(),
               attributes.maximumQuantity(), attributes.minimumQuantity(), attributes.chance());
      } catch (JsonProcessingException e) {
         System.out.printf("Unable to process line [%s]", line);
      }
      return null;
   }

   public static List<MonsterDropData> getAll() {
      return Connection.instance()
            .list(MonsterDropProvider::getAll);
   }

   public static List<MonsterDropData> getByMonsterId(int monsterId) {
      return Connection.instance()
            .list(entityManager -> MonsterDropProvider.getByMonsterId(entityManager, monsterId));
   }
}
