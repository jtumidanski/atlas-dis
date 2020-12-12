package com.atlas.dis.rest.processor;

import builder.ResultBuilder;
import com.app.rest.util.stream.Collectors;
import com.atlas.dis.rest.ResultObjectFactory;

import javax.ws.rs.core.Response;

public final class MonsterDropProcessor {
   private MonsterDropProcessor() {
   }

   public static ResultBuilder getAll() {
      return com.atlas.dis.processor.MonsterDropProcessor.getAll().stream()
            .map(ResultObjectFactory::create)
            .collect(Collectors.toResultBuilder());
   }

   public static ResultBuilder getByMonsterId(int monsterId) {
      return com.atlas.dis.processor.MonsterDropProcessor.getByMonsterId(monsterId).stream()
            .map(ResultObjectFactory::create)
            .collect(Collectors.toResultBuilder());
   }
}
