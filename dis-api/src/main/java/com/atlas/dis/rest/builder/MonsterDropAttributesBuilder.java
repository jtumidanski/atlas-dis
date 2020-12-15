package com.atlas.dis.rest.builder;

import builder.AttributeResultBuilder;
import com.app.common.builder.RecordBuilder;
import com.atlas.dis.rest.attribute.MonsterDropAttributes;

public class MonsterDropAttributesBuilder extends RecordBuilder<MonsterDropAttributes, MonsterDropAttributesBuilder> implements AttributeResultBuilder {
   private Integer monsterId;

   private Integer itemId;

   private Integer maximumQuantity;

   private Integer minimumQuantity;

   private Integer chance;


   @Override
   public MonsterDropAttributes construct() {
      return new MonsterDropAttributes(monsterId, itemId, maximumQuantity, minimumQuantity, chance);
   }

   @Override
   public MonsterDropAttributesBuilder getThis() {
      return this;
   }

   public MonsterDropAttributesBuilder setMonsterId(Integer monsterId) {
      this.monsterId = monsterId;
      return getThis();
   }

   public MonsterDropAttributesBuilder setItemId(Integer itemId) {
      this.itemId = itemId;
      return getThis();
   }

   public MonsterDropAttributesBuilder setMaximumQuantity(Integer maximumQuantity) {
      this.maximumQuantity = maximumQuantity;
      return getThis();
   }

   public MonsterDropAttributesBuilder setMinimumQuantity(Integer minimumQuantity) {
      this.minimumQuantity = minimumQuantity;
      return getThis();
   }

   public MonsterDropAttributesBuilder setChance(Integer chance) {
      this.chance = chance;
      return getThis();
   }
}
