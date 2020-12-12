package com.atlas.dis.rest;

import builder.ResultObjectBuilder;
import com.atlas.dis.model.MonsterDropData;
import com.atlas.dis.rest.attribute.MonsterDropAttributes;
import com.atlas.dis.rest.builder.MonsterDropAttributesBuilder;

public final class ResultObjectFactory {
   private ResultObjectFactory() {
   }

   public static ResultObjectBuilder create(MonsterDropData monsterDropData) {
      return new ResultObjectBuilder(MonsterDropAttributes.class, monsterDropData.id())
            .setAttribute(new MonsterDropAttributesBuilder()
                  .setMonsterId(monsterDropData.monsterId())
                  .setItemId(monsterDropData.itemId())
                  .setMaximumQuantity(monsterDropData.maximumQuantity())
                  .setMinimumQuantity(monsterDropData.minimumQuantity())
                  .setChance(monsterDropData.chance())
            );
   }
}
